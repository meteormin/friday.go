// 공통 타입 정의

export interface Token {
    token: string
    expiresAt: string
    issuedAt: string
}

export interface WithToken {
    getToken: () => Token,
    setToken: (token: Token) => void
}
