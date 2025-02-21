// 공통 타입 정의

export interface Token {
    exp: number
    token: string
}

export interface WithToken {
    getToken: () => Token,
    setToken: (token: Token) => void
}
