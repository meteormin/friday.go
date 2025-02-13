import React from 'react';
import CssBaseline from '@mui/material/CssBaseline';
import Container from '@mui/material/Container';
import AppTheme from './theme/AppTheme';
import MainAppBar, {MainAppBarProps} from './MainAppBar';
import Footer from './Footer';

export interface MainContainerProps {
    disableCustomTheme?: boolean;
    mainAppBarProps: MainAppBarProps
    children?: React.ReactNode;
}

export default function MainContainer(props: MainContainerProps) {
    return (
        <AppTheme disableCustomTheme={props.disableCustomTheme}>
            <CssBaseline enableColorScheme/>
            <MainAppBar title={props.mainAppBarProps?.title ?? ''}
                        navItems={props.mainAppBarProps?.navItems ?? []}
                        isLogin={props.mainAppBarProps?.isLogin ?? false}
            />
            <Container
                maxWidth="lg"
                component="main"
                sx={{display: 'flex', flexDirection: 'column', my: 16, gap: 4}}
            >
                {props.children}
            </Container>
            <Footer/>
        </AppTheme>
    );
}