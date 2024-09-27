package service_test

import (
	"context"
	"errors"
	"html/template"
	"strconv"
	"testing"
	"time"

	"github.com/hexley21/fixup/internal/user/delivery/http/v1/dto"
	"github.com/hexley21/fixup/internal/user/enum"
	"github.com/hexley21/fixup/internal/user/repository"
	mock_repository "github.com/hexley21/fixup/internal/user/repository/mock"
	"github.com/hexley21/fixup/internal/user/service"
	mock_encryption "github.com/hexley21/fixup/pkg/encryption/mock"
	"github.com/hexley21/fixup/pkg/hasher"
	mock_hasher "github.com/hexley21/fixup/pkg/hasher/mock"
	mock_cdn "github.com/hexley21/fixup/pkg/infra/cdn/mock"
	mock_postgres "github.com/hexley21/fixup/pkg/infra/postgres/mock"
	mock_mailer "github.com/hexley21/fixup/pkg/mailer/mock"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var (
	registerUserDto = dto.RegisterUser{
		Email:       "larry@page.com",
		PhoneNumber: "995111222333",
		FirstName:   "Larry",
		LastName:    "Page",
		Password:    "12345678",
	}

	registerProviderDto = dto.RegisterProvider{
		RegisterUser: registerUserDto,
		PersonalIDNumber: "1234567890",
	}

	loginDto = dto.Login{
		Email: "larry@page.com",
		Password: "12345678",
	}

	creds = repository.GetCredentialsByEmailRow{
		ID: 1,
		Role: enum.UserRoleADMIN,
		Hash: newHash,
		UserStatus: pgtype.Bool{Bool: true, Valid: true},
	}

	emailAddress = "fixup@gmail.com"
	newHash          = "hash"
	newUrl           = "picture.png?signed=true"

	confiramtionTemplate = template.New("confirmation")
	verifiedTemplate     = template.New("verified")

	verificationTTL = time.Hour
)

func setupAuth(t *testing.T) (
	ctrl *gomock.Controller,
	svc service.AuthService,
	mockUserRepo *mock_repository.MockUserRepository,
	mockProviderRepo *mock_repository.MockProviderRepository,
	mockVerificationRepo *mock_repository.MockVerificationRepository,
	mockPgx *mock_postgres.MockPGX,
	mockTx *mock_postgres.MockTx,
	mockHasher *mock_hasher.MockHasher,
	mockEncryptor *mock_encryption.MockEncryptor,
	mockMailer *mock_mailer.MockMailer,
	mockUrlSigner *mock_cdn.MockURLSigner,
) {
	ctrl = gomock.NewController(t)

	mockUserRepo = mock_repository.NewMockUserRepository(ctrl)
	mockProviderRepo = mock_repository.NewMockProviderRepository(ctrl)
	mockVerificationRepo = mock_repository.NewMockVerificationRepository(ctrl)
	mockPgx = mock_postgres.NewMockPGX(ctrl)
	mockTx = mock_postgres.NewMockTx(ctrl)
	mockHasher = mock_hasher.NewMockHasher(ctrl)
	mockEncryptor = mock_encryption.NewMockEncryptor(ctrl)
	mockMailer = mock_mailer.NewMockMailer(ctrl)
	mockUrlSigner = mock_cdn.NewMockURLSigner(ctrl)

	s := service.NewAuthService(mockUserRepo, mockProviderRepo, mockVerificationRepo, verificationTTL, mockPgx, mockHasher, mockEncryptor, mockMailer, emailAddress, mockUrlSigner)
	s.SetTemplates(confiramtionTemplate, verifiedTemplate)

	svc = s
	return
}

func TestRegisterCustomer(t *testing.T) {
	ctx := context.Background()

	ctrl, service, mockUserRepo, _, _, _, _, mockHasher, _, _, mockUrlSigner := setupAuth(t)
	defer ctrl.Finish()

	mockUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return(userEntity, nil)
	mockHasher.EXPECT().HashPassword(gomock.Any()).Return(newHash)
	mockUrlSigner.EXPECT().SignURL(gomock.Any()).Return(newUrl, nil)

	dto, err := service.RegisterCustomer(ctx, registerUserDto)
	assert.NoError(t, err)
	assert.Equal(t, registerUserDto.Email, dto.Email)
	assert.Equal(t, registerUserDto.PhoneNumber, dto.PhoneNumber)
	assert.Equal(t, registerUserDto.FirstName, dto.FirstName)
	assert.Equal(t, registerUserDto.LastName, dto.LastName)

	time.Sleep(time.Microsecond)
}

func TestRegisterProvider(t *testing.T) {
	ctx := context.Background()

	ctrl, service, mockUserRepo, mockProviderRepo, _, mockPgx, mockTx, mockHasher, mockEncryptor, _, mockUrlSigner := setupAuth(t)
	defer ctrl.Finish()

	mockUserRepo.EXPECT().WithTx(mockTx).Return(mockUserRepo)
	mockUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Return(userEntity, nil)
	mockProviderRepo.EXPECT().WithTx(mockTx).Return(mockProviderRepo)
	mockProviderRepo.EXPECT().Create(ctx, gomock.Any()).Return(nil)

	mockPgx.EXPECT().BeginTx(ctx, gomock.Any()).Return(mockTx, nil)
	mockTx.EXPECT().Commit(ctx).Return(nil)

	mockHasher.EXPECT().HashPassword(gomock.Any()).Return(newHash)
	mockEncryptor.EXPECT().Encrypt(gomock.Any()).Return([]byte(registerProviderDto.PersonalIDNumber), nil)
	mockUrlSigner.EXPECT().SignURL(gomock.Any()).Return(newUrl, nil)

	dto, err := service.RegisterProvider(ctx, registerProviderDto)
	assert.NoError(t, err)
	assert.Equal(t, registerUserDto.Email, dto.Email)
	assert.Equal(t, registerUserDto.PhoneNumber, dto.PhoneNumber)
	assert.Equal(t, registerUserDto.FirstName, dto.FirstName)
	assert.Equal(t, registerUserDto.LastName, dto.LastName)

	time.Sleep(time.Microsecond)
}

func TestAuthenticateUser_Success(t *testing.T) {
	ctx := context.Background()

	ctrl, _, mockUserRepo, _, _, _, _, mockHasher, _, _, _ := setupAuth(t)
	defer ctrl.Finish()

	mockUserRepo.EXPECT().GetCredentialsByEmail(ctx, loginDto.Email).Return(creds, nil)
	mockHasher.EXPECT().VerifyPassword(loginDto.Password, creds.Hash).Return(nil)

	service := service.NewAuthService(mockUserRepo, nil, nil, time.Hour, nil, mockHasher, nil, nil, emailAddress, nil)

	credentialsDto, err := service.AuthenticateUser(ctx, loginDto)
	assert.NoError(t, err)
	assert.Equal(t, strconv.FormatInt(creds.ID, 10), credentialsDto.ID)
	assert.Equal(t, string(credentialsDto.Role), credentialsDto.Role)
	assert.Equal(t, creds.UserStatus.Bool, credentialsDto.UserStatus)
}

func TestAuthenticateUser_NotFound(t *testing.T) {
	ctx := context.Background()

	ctrl, svc, mockUserRepo, _, _, _, _, _, _, _, _ := setupAuth(t)
	defer ctrl.Finish()

	mockUserRepo.EXPECT().GetCredentialsByEmail(ctx, loginDto.Email).Return(repository.GetCredentialsByEmailRow{}, pgx.ErrNoRows)

	credentialsDto, err := svc.AuthenticateUser(ctx, loginDto)
	assert.ErrorIs(t, err, pgx.ErrNoRows)
	assert.Empty(t, credentialsDto)
}

func TestAuthenticateUser_PasswordMissmatch(t *testing.T) {
	ctx := context.Background()

	ctrl, svc, mockUserRepo, _, _, _, _, mockHasher, _, _, _ := setupAuth(t)
	defer ctrl.Finish()

	mockUserRepo.EXPECT().GetCredentialsByEmail(ctx, loginDto.Email).Return(creds, nil)
	mockHasher.EXPECT().VerifyPassword(loginDto.Password, creds.Hash).Return(hasher.ErrPasswordMismatch)

	credentialsDto, err := svc.AuthenticateUser(ctx, loginDto)
	assert.ErrorIs(t, err, hasher.ErrPasswordMismatch)
	assert.Empty(t, credentialsDto)
}

func TestVerifyUser(t *testing.T) {
	ctx := context.Background()

	ctrl, svc, mockUserRepo, _, _, _, _, _, _, _, _ := setupAuth(t)
	defer ctrl.Finish()

	mockUserRepo.EXPECT().UpdateStatus(ctx, gomock.Any()).Return(nil)

	assert.NoError(t, svc.VerifyUser(ctx, 1, ""))

	time.Sleep(100 * time.Millisecond)
}

func TestGetUserConfirmationDetails_Success(t *testing.T) {
	ctx := context.Background()

	ctrl, svc, mockUserRepo, _, _, _, _, _, _, _, _ := setupAuth(t)
	defer ctrl.Finish()

	args := repository.GetUserConfirmationDetailsRow{
		ID: 1,
		UserStatus: pgtype.Bool{Bool: false, Valid: true},
		FirstName: "Larry",
	}

	mockUserRepo.EXPECT().GetUserConfirmationDetails(ctx, gomock.Any()).Return(args, nil)

	dto, err := svc.GetUserConfirmationDetails(ctx, "")
	assert.NoError(t, err)
	assert.Equal(t, strconv.FormatInt(args.ID, 10), dto.ID)
	assert.Equal(t, args.UserStatus.Bool, dto.UserStatus)
	assert.Equal(t, args.FirstName, dto.Firstname)
}

func TestSendConfirmationLetter_Success(t *testing.T) {
	ctx := context.Background()

	ctrl, svc, _, _, mockVerificationRepo, _, _, _, _, mockMailer, _ := setupAuth(t)
	defer ctrl.Finish()

	mockVerificationRepo.EXPECT().SetTokenUsed(ctx, gomock.Any(), verificationTTL)
	mockMailer.EXPECT().SendHTML(emailAddress, gomock.Any(), gomock.Any(), confiramtionTemplate, gomock.Any()).Return(nil)

	assert.NoError(t, svc.SendConfirmationLetter(ctx, "", "", ""))
}

func TestSendConfirmationLetter_AlreadySet(t *testing.T) {
	ctx := context.Background()

	ctrl, svc, _, _, mockVerificationRepo, _, _, _, _, _, _ := setupAuth(t)
	defer ctrl.Finish()

	mockVerificationRepo.EXPECT().SetTokenUsed(ctx, gomock.Any(), verificationTTL).Return(redis.TxFailedErr)

	assert.ErrorIs(t, redis.TxFailedErr, svc.SendConfirmationLetter(ctx, "", "", ""))
}

func TestSendVerifiedLetter_Success(t *testing.T) {
	ctrl, svc, _, _, _, _, _, _, _, mockMailer, _ := setupAuth(t)
	defer ctrl.Finish()

	mockMailer.EXPECT().SendHTML(emailAddress, gomock.Any(), gomock.Any(), verifiedTemplate, gomock.Nil()).Return(nil)

	assert.NoError(t, svc.SendVerifiedLetter(""))
}


func TestIsVerificationTokenUsed_Used(t *testing.T) {
	ctx := context.Background()

	ctrl, svc, _, _, mockVerificationRepo, _, _, _, _, _, _ := setupAuth(t)
	defer ctrl.Finish()

	mockVerificationRepo.EXPECT().IsTokenUsed(ctx, gomock.Any()).Return(true, nil)

	isUsed, err := svc.IsVerificationTokenUsed(ctx, "")
	assert.True(t, isUsed)
	assert.NoError(t,err)
}

func TestIsVerificationTokenUsed_NotUsed(t *testing.T) {
	ctx := context.Background()

	ctrl, svc, _, _, mockVerificationRepo, _, _, _, _, _, _ := setupAuth(t)
	defer ctrl.Finish()

	mockVerificationRepo.EXPECT().IsTokenUsed(ctx, gomock.Any()).Return(false, nil)

	isUsed, err := svc.IsVerificationTokenUsed(ctx, "")
	assert.False(t, isUsed)
	assert.NoError(t,err)
}

func TestIsVerificationTokenUsed_RepoError(t *testing.T) {
	ctx := context.Background()

	ctrl, svc, _, _, mockVerificationRepo, _, _, _, _, _, _ := setupAuth(t)
	defer ctrl.Finish()

	mockVerificationRepo.EXPECT().IsTokenUsed(ctx, gomock.Any()).Return(false, errors.New(""))

	isUsed, err := svc.IsVerificationTokenUsed(ctx, "")
	assert.False(t, isUsed)
	assert.Error(t,err)
}