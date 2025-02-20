import React, {useState} from "react";
import {Box, Button, Container, TextField, Typography} from "@mui/material";
import * as axios from "axios";
import config from "../../config.ts"

export default function SignIn() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const handleLogin = () => {
        console.log("Email:", email);
        console.log("Password:", password);
        axios.default.post(config.apiUrl)
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
                    로그인
                </Typography>
                <TextField
                    label="이메일"
                    variant="outlined"
                    fullWidth
                    margin="normal"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
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
                    onClick={handleLogin}
                >
                    로그인
                </Button>
            </Box>
        </Container>
    );
};
