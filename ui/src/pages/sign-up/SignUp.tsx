import React, {useState} from "react";
import {Box, Button, Container, createTheme, CssBaseline, TextField, ThemeProvider, Typography} from "@mui/material";
import {AuthClient} from "../../api/auth";
import Grid from "@mui/material/Grid2";
import {Link} from "react-router-dom";


export interface SignUpProps {
    authClient: AuthClient
}

interface SignUpForm {
    name: string;
    username: string;
    password: string;
}

export default function SignUp(props: SignUpProps) {

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

    const [signUpForm, setSignUpForm] = useState<SignUpForm>({
        name: "",
        username: "",
        password: "",
    });

    const [errors, setErrors] = useState({
        name: "",
        username: "",
        password: "",
    });

    const handleRegister = () => {
        let isValid = true;
        if (signUpForm.name === "") {
            setErrors((prev) => ({
                ...prev,
                name: "이름을 입력하세요."
            }));
            isValid = false;
        } else {
            setErrors((prev) => ({
                ...prev,
                name: ""
            }));
        }

        if (signUpForm.username === "") {
            setErrors((prev) => ({
                ...prev,
                username: "로그인 ID를 입력하세요."
            }));
            isValid = false;
        } else {
            setErrors((prev) => ({
                ...prev,
                username: ""
            }));
        }

        if (signUpForm.password === "") {
            setErrors((prev) => ({
                ...prev,
                password: "비밀번호를 입력하세요."
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

        props.authClient.signUp({
            name: signUpForm.name,
            username: signUpForm.username,
            password: signUpForm.password
        }).then(() => {
            alert("회원가입 성공");
            window.location.href = "/sign-in";
        }).catch((error: any) => {
            alert("회원가입 실패");
            console.error(error);
        })
    };

    const setName = (name: string) => {
        setSignUpForm((prev) => ({
            ...prev,
            name: name
        }));
    };

    const setUsername = (username: string) => {
        setSignUpForm((prev) => ({
            ...prev,
            username: username
        }));
    };

    const setPassword = (password: string) => {
        setSignUpForm((prev) => ({
            ...prev,
            password: password
        }));
    };

    return (
        <ThemeProvider theme={theme}>
            <CssBaseline/>
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
                        회원가입
                    </Typography>
                    <TextField
                        label="이름"
                        variant="outlined"
                        required
                        fullWidth
                        margin="normal"
                        value={signUpForm.name}
                        error={!!errors.name}
                        helperText={errors.name}
                        onChange={(e) => setName(e.target.value)}
                    />
                    <TextField
                        label="로그인 ID"
                        variant="outlined"
                        required
                        fullWidth
                        margin="normal"
                        value={signUpForm.username}
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
                        value={signUpForm.password}
                        error={!!errors.password}
                        helperText={errors.password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    <Button
                        variant="contained"
                        color="primary"
                        fullWidth
                        sx={{mt: 2}}
                        onClick={handleRegister}
                    >
                        회원가입
                    </Button>

                    <Grid container justifyContent="flex-end" sx={{width: "100%", mt: 1}}>
                        <Grid>
                            <Button component={Link} to="/sign-in" color="secondary" sx={{textTransform: "none"}}>
                                로그인
                            </Button>
                        </Grid>
                    </Grid>
                </Box>
            </Container>
        </ThemeProvider>
    );
};