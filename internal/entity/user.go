package entity

type RegisterReq struct {
    Username     string `json:"username"`
    Email        string `json:"email"`
    Password     string `json:"password"`
    FullName     string `json:"full_name"`
    DateOfBirth  string `json:"date_of_birth"`
}

type User struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Role     string `json:"role"`
}

type LoginReq struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginRes struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
    Role         string `json:"role"`
}

type GetByEmail struct {
    Email string `json:"email"`
}

type ResetPassReq struct {
    ResetToken  string `json:"reset_token"`
    Email       string `json:"email"`
    NewPassword string `json:"new_password"`
}

type ResetPassReqBody struct {
    ResetToken  string `json:"reset_token"`
    NewPassword string `json:"new_password"`
}

type Params struct {
    From    string `json:"from"`
    Password string `json:"password"`
    To      string `json:"to"`
    Message string `json:"message"`
    Code    string `json:"code"`
}

type RefToken struct {
    ID        string `json:"id"`
    UserID    string `json:"user_id"`
    Token     string `json:"token"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    DeletedAt string `json:"deleted_at"`
}

type ListUserReq struct {
    Username string         `json:"username"`
    FullName string         `json:"full_name"`
    Filter   *Filter `json:"filter"` 
}

type ListUserRes struct {
    Users []*UserRes `json:"users"`
    Count int32     `json:"count"`
}

type UserRes struct {
    ID          string `json:"id"`
    Username    string `json:"username"`
    Email       string `json:"email"`
    FullName    string `json:"full_name"`
    DateOfBirth string `json:"date_of_birth"`
    Role        string `json:"role"`
}

type EditProfileReqBody struct {
    Username    string `json:"username"`
    Email       string `json:"email"`
    FullName    string `json:"full_name"`
    DateOfBirth string `json:"date_of_birth"`
}

type ChangePasswordReq struct {
    ID             string `json:"id"`
    CurrentPassword string `json:"current_password"`
    NewPassword    string `json:"new_password"`
}

type ChangePasswordReqBody struct {
    CurrentPassword string `json:"current_password"`
    NewPassword     string `json:"new_password"`
}

type SettingReq struct {
    ID            string `json:"id"`
    PrivacyLevel  string `json:"privacy_level"`
    Notification  string `json:"notification"`
    Language      string `json:"language"`
    Theme         string `json:"theme"`
}

type Setting struct {
    PrivacyLevel  string `json:"privacy_level"`
    Notification  string `json:"notification"`
    Language      string `json:"language"`
    Theme         string `json:"theme"`
}

type Void struct{}

type GetById struct {
    ID string `json:"id"`
}

type Filter struct {
    Limit  int32 `json:"limit"`
    Offset int32 `json:"offset"`
}
