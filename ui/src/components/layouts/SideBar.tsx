import {
    Box,
    Drawer,
    List,
    ListItem,
    ListItemButton,
    ListItemIcon,
    ListItemText,
    Toolbar,
    Typography
} from "@mui/material";
import {Link} from "react-router-dom";
import React from "react";

export interface MenuItem {
    text: string;
    icon: React.ReactNode;
    href: string;
}

export interface SideBarProps {
    appName : string;
    open: boolean;
    items: MenuItem[];
}

export default function Sidebar(props: SideBarProps) {
    return (
        <Drawer
            variant="permanent"
            sx={{
                width: props.open ? 240 : 60, // ✅ 닫힌 상태일 때 최소 크기 유지
                flexShrink: 0,
                [`& .MuiDrawer-paper`]: {
                    width: props.open ? 240 : 60,
                    boxSizing: "border-box",
                },
            }}
        >
            {/* ✅ Toolbar 내부에 My App 추가 (사이드바 열렸을 때만 표시) */}
            <Toolbar>
                {props.open && (
                    <Typography variant="h6" sx={{textAlign: "left", width: "100%", justifyContent: "center"}}>
                        {props.appName}
                    </Typography>
                )}
            </Toolbar>

            <List>
                {props.items.map((item) => (
                    <ListItem key={item.text} disablePadding sx={{ paddingY: 0.5 }}>
                        <ListItemButton component={Link} to={item.href}>
                            <ListItemIcon>{item.icon}</ListItemIcon>
                            {/* ✅ 사이드바 닫힐 때 텍스트 숨김 */}
                            <Box sx={{display: props.open ? "block" : "none", transition: "display 0.3s ease"}}>
                                <ListItemText primary={item.text}/>
                            </Box>
                        </ListItemButton>
                    </ListItem>
                ))}
            </List>
        </Drawer>
    );
}