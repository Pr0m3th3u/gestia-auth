package handlers

import "net/http"

type AuthDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Credentials struct {
	AccessToken   string
	RefreshTokent string
}

type UserDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type IRootHandler interface {
	// LoginHandler(authDto AuthDTO) (Credentials, error)
	// RefreshTokenHandler(credentials Credentials) (Credentials, error)
	// RegisterHandler(userDTO UserDTO) error
	LoginHandler(w http.ResponseWriter, r *http.Request)
	RefreshTokenHandler(w http.ResponseWriter, r *http.Request)
	RegisterHandler(w http.ResponseWriter, r *http.Request)
}

type RootHandler struct {
}

var (
	_ IRootHandler = (*RootHandler)(nil)
)

func NewRootHandler() IRootHandler {
	return &RootHandler{}
}

// LoginHandler выполняет авторизацию пользователя
// @Summary Авторизация пользователя
// @Description Аутентификация пользователя с возвратом access и refresh токенов
// @Tags auth
// @Accept json
// @Produce json
// @Param input body AuthDTO true "Данные для авторизации"
// @Success 200 {object} Credentials
// @Failure 400 {object} map[string]string "Ошибка валидации"
// @Failure 401 {object} map[string]string "Неверный логин или пароль"
// @Router /api/v1/auth/login [post]
func (rh *RootHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// RefreshTokenHandler обновляет токен доступа
// @Summary Обновление access токена
// @Description Используется для получения нового access токена по refresh токену
// @Tags auth
// @Accept json
// @Produce json
// @Param input body Credentials true "Refresh токен"
// @Success 200 {object} Credentials
// @Failure 400 {object} map[string]string "Ошибка валидации"
// @Failure 401 {object} map[string]string "Неверный refresh токен"
// @Router /api/v1/auth/refresh [post]
func (rh *RootHandler) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// RegisterHandler выполняет регистрацию нового пользователя
// @Summary Регистрация нового пользователя
// @Description Создание нового пользователя с указанными данными
// @Tags auth
// @Accept json
// @Produce json
// @Param input body UserDTO true "Данные для регистрации"
// @Success 201 {string} string "Пользователь успешно зарегистрирован"
// @Failure 400 {object} map[string]string "Ошибка валидации"
// @Failure 409 {object} map[string]string "Пользователь уже существует"
// @Router /api/v1/auth/register [post]
func (rh *RootHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}
