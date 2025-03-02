import {AppBar, Box, createTheme, CssBaseline, IconButton, ThemeProvider, Toolbar} from "@mui/material";
import Sidebar, {MenuItem} from "./layouts/SideBar";
import React, {useEffect, useState} from "react";
import Brightness4Icon from '@mui/icons-material/Brightness4';
import Brightness7Icon from '@mui/icons-material/Brightness7';
import MenuIcon from "@mui/icons-material/Menu";
import HomeIcon from "@mui/icons-material/Home";
import WebIcon from "@mui/icons-material/Web";
import TagIcon from "@mui/icons-material/Tag";
import HistoryIcon from "@mui/icons-material/History";
import LogoutIcon from '@mui/icons-material/Logout';

const defaultMenuItems: MenuItem[] = [
    {text: "홈", icon: <HomeIcon/>, href: "/"},
    {text: "사이트", icon: <WebIcon/>, href: "/sites"},
    {text: "태그", icon: <TagIcon/>, href: "/tags"},
    {text: "기록", icon: <HistoryIcon/>, href: "/history"},
    {text: "로그아웃", icon: <LogoutIcon sx={{marginLeft: "3px"}}/>, href: "/logout"},
];

export default function LayoutContainer(props: { appName:string, menuItems?: MenuItem[], children: React.ReactNode }) {
    // ✅ 로컬 스토리지에서 다크 모드 상태 불러오기
    const [darkMode, setDarkMode] = useState(() => {
        const saved = localStorage.getItem("darkMode");
        return saved === "true"; // 기본값: 라이트 모드
    });

    // ✅ 로컬 스토리지에서 사이드바 상태 불러오기
    const [sidebarOpen, setSidebarOpen] = useState(() => {
        const saved = localStorage.getItem("sidebarOpen");
        return saved !== "false"; // 기본값: 열림 (true)
    });

    // ✅ 사이드바 상태 변경 시 `localStorage`에 저장
    useEffect(() => {
        localStorage.setItem("sidebarOpen", sidebarOpen.toString());
        localStorage.setItem("darkMode", darkMode.toString());
    }, [sidebarOpen, darkMode]);

    // ✅ 사이드바 토글 함수
    const toggleSidebar = () => {
        setSidebarOpen((prev) => !prev);
    };
    // ✅ MUI 다크/라이트 모드 테마 정의
    const theme = createTheme({
        palette: {
            mode: darkMode ? "dark" : "light",
        },
    });

    const toggleDarkMode = () => {
        setDarkMode(!darkMode);
    }

    return (
        <ThemeProvider theme={theme}>
            <CssBaseline/> {/* ✅ 기본 스타일 적용 (배경색 변경 등) */}
            <Box sx={{display: "flex"}}>
                {/* ✅ 상단 네비게이션 바 (AppBar) */}
                <AppBar
                    position="fixed"
                    sx={{
                        zIndex: (theme) => theme.zIndex.drawer + 1, // ✅ Sidebar 위에 오도록 설정
                        width: `calc(100% - ${sidebarOpen ? 240 : 60}px)`, // ✅ Sidebar 크기에 맞춰 조정
                        marginLeft: `${sidebarOpen ? 240 : 60}px`, // ✅ Sidebar가 닫히면 margin 조정
                        transition: "width 0.3s ease, margin-left 0.3s ease",
                    }}
                >
                    <Toolbar sx={{display: "flex", justifyContent: "space-between"}}>
                        {/* ✅ 햄버거 버튼 (사이드바 토글) */}
                        <IconButton onClick={toggleSidebar} color="inherit">
                            <MenuIcon/>
                        </IconButton>

                        {/* ✅ 다크 모드 토글 버튼 (우측 정렬) */}
                        <IconButton onClick={toggleDarkMode} color="inherit">
                            {darkMode ? <Brightness7Icon/> : <Brightness4Icon/>}
                        </IconButton>
                    </Toolbar>
                </AppBar>

                {/* ✅ 좌측 사이드바 */}
                <Sidebar appName={props.appName} items={props.menuItems ?? defaultMenuItems} open={sidebarOpen}/>

                {/* ✅ 메인 콘텐츠 */}
                <Box component="main" sx={{flexGrow: 1, p: 3, mt: 8}}>
                    {props.children}
                </Box>
            </Box>
        </ThemeProvider>
    );
};