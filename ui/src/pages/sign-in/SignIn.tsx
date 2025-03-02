import React, {useState} from "react";
import {Link} from "react-router-dom";
import {Box, Button, Container, createTheme, CssBaseline, TextField, ThemeProvider, Typography} from "@mui/material";
import {AuthClient} from "../../api/auth";
import Grid from "@mui/material/Grid2";

export interface SignInProps {
    authClient: AuthClient
}

interface SignInForm {
    username: string;
    password: string;
}

export default function SignIn(props: SignInProps) {
    // ✅ 로컬 스토리지에서 다크 모드 상태 불러오기
    const getDarkMode = () => {
        const saved = localStorage.getItem("darkMode");
        return saved === "true"; // 기본값: 라이트 모드
    }

    // ✅ MUI 다크/라이트 모드 테마 정의
    const theme = createTheme({
        palette: {
            mode: getDarkMode() ? "dark" : "light",
        },
    });

    const [signInForm, setSignInForm] = useState<SignInForm>({
        username: "",
        password: "",
    });

    const [errors, setErrors] = useState({
        username: "",
        password: "",
    });

    const handleLogin = () => {
        let isValid = true;
        if (signInForm.username === "") {
            setErrors((prev) => ({
                ...prev,
                username: "로그인 ID를 입력해주세요."
            }));
            isValid = false;
        } else {
            setErrors((prev) => ({
                ...prev,
                username: ""
            }));
        }

        if (signInForm.password === "") {
            setErrors((prev) => ({
                ...prev,
                password: "비밀번호를 입력해주세요."
            }));
            isValid = false;
        } else {
            setErrors((prev) => ({
                ...prev,
                password: ""
            }));
        }

        if (!isValid) {
            return;
        }

        props.authClient.signIn({
            username: signInForm.username,
            password: signInForm.password
        }).then((res) => {
            alert("로그인 성공");

            props.authClient.setToken({
                token: res.token,
                expiresAt: res.expiresAt,
                issuedAt: res.issuedAt
            });

            window.location.href = "/";
        }).catch((error: any) => {
            alert("로그인 실패");
            console.error(error);
        });
    };

    const setUsername = (username: string) => {
        setSignInForm((prev) => ({
            ...prev,
            username: username
        }));
    };

    const setPassword = (password: string) => {
        setSignInForm((prev) => ({
            ...prev,
            password: password
        }));
    };

    return (
        <ThemeProvider theme={theme}>
            <CssBaseline/> {/* ✅ 기본 스타일 적용 (배경색 변경 등) */}
            <Container maxWidth="xs">
                <Box
                    sx={{
                        display: "flex",
                        flexDirection: "column",
                        alignItems: "center",
                        justifyContent: "center",
                        height: "100vh",
                    }}
                >
                    <Typography variant="h5" gutterBottom>
                        로그인
                    </Typography>
                    <TextField
                        label="로그인 ID"
                        variant="outlined"
                        required
                        fullWidth
                        margin="normal"
                        value={signInForm.username}
                        error={!!errors.username}
                        helperText={errors.username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                    <TextField
                        label="비밀번호"
                        type="password"
                        variant="outlined"
                        required
                        fullWidth
                        margin="normal"
                        value={signInForm.password}
                        error={!!errors.password}
                        helperText={errors.password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    <Button
                        variant="contained"
                        color="primary"
                        fullWidth
                        sx={{mt: 2}}
                        onClick={handleLogin}
                    >
                        로그인
                    </Button>

                    <Grid container justifyContent="flex-end" sx={{width: "100%", mt: 1}}>
                        <Grid>
                            <Button component={Link} to="/sign-up" color="secondary" sx={{textTransform: "none"}}>
                                회원가입
                            </Button>
                        </Grid>
                    </Grid>
                </Box>
            </Container>
        </ThemeProvider>
    );
};
