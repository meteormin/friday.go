import {BrowserRouter, Route, Routes} from "react-router";
import {Navigate} from "react-router-dom";
import * as React from "react";
import SignIn from "./pages/sign-in/SignIn";
import SignUp from "./pages/sign-up/SignUp";
import Posts from "./pages/posts/Posts.tsx";
import {ApiClient, newApiClient} from "./api";
import config from "./config.ts";
import {Token} from "./api/common.ts";
import LayoutContainer from "./components/LayoutContainer.tsx";
import dayjs from "dayjs";

function Guard({getToken, children}: { getToken: () => Token, children: React.ReactNode }) {
    const token = getToken();
    const exp = dayjs(token.expiresAt, 'YYYY-MM-DD HH:mm:ss');
    const now = dayjs();
    if (now.isBefore(exp)) {
        return <LayoutContainer appName={config.appName}>
            {children}
        </LayoutContainer>
    } else {
        return <Navigate to="/sign-in" replace/>
    }
}

function Logout() {
    localStorage.removeItem("token");
    return <Navigate to="/sign-in" replace/>
}

interface RouteProps {
    path: string,
    element: React.ReactNode
}

const routes = (): RouteProps[] => {
    const apiClient: ApiClient = newApiClient(config.apiUrl);

    return [
        {
            path: "/",
            element: <Guard getToken={apiClient.getToken}><Posts postsClient={apiClient.posts}/></Guard>
        },
        {
            path: "/posts",
            element: <Guard getToken={apiClient.getToken}><Posts postsClient={apiClient.posts}/></Guard>
        },
        {
            path: "/sign-in",
            element: <SignIn authClient={apiClient.auth}/>
        },
        {
            path: "/sign-up",
            element: <SignUp authClient={apiClient.auth}/>
        },
        {
            path: "/logout",
            element: <Guard getToken={apiClient.getToken}>
                <Logout/>
            </Guard>
        }
    ];
}

function App() {
    return (
        <BrowserRouter>
            <Routes>
                {routes().map((route) => (
                    <Route key={route.path} path={route.path} element={route.element}/>
                ))}
            </Routes>
        </BrowserRouter>
    )
}

export default App
