namespace go user

struct BaseResp {
    1: i64 code,
    2: string msg,
}

struct User {
    1: i64 id,
    2: string name,
}

struct RegisterRequest {
    1: string username,
    2: string password,
}

struct RegisterResponse {
    1: BaseResp base,
    2: i64 user_id,
    3: string token,
}

struct LoginRequest {
    1: string username,
    2: string password,
}

struct LoginResponse {
    1: BaseResp base,
    2: User user,
    3: string token,
}

struct InfoRequest {
    1: i64 user_id,
    2: string token,
}

struct InfoResponse {
    1: BaseResp base,
    2: User user,
}

struct ChangeRequest {
    1: string username,
    2: string password,
    3: string new_password,
}

struct ChangeResponse {
    1: BaseResp base,
}

service UserService {
    RegisterResponse Register(1: RegisterRequest req),
    LoginResponse Login(1: LoginRequest req),
    InfoResponse Info(1: InfoRequest req),
    ChangeResponse Change(1: ChangeRequest req),
}
