import React, {useState} from "react";
import {Box, Button, Container, TextField, Typography} from "@mui/material";

export default function SignUp() {
    const [name, setName] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const handleRegister = () => {
        console.log("Name:", name);
        console.log("Username:", username);
        console.log("Password:", password);
    };

    return (
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
                    fullWidth
                    margin="normal"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                />
                <TextField
                    label="로그인 ID(이메일)"
                    variant="outlined"
                    fullWidth
                    margin="normal"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                />
                <TextField
                    label="비밀번호"
                    type="password"
                    variant="outlined"
                    fullWidth
                    margin="normal"
                    value={password}
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
            </Box>
        </Container>
    );
};