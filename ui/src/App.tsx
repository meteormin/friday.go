import {BrowserRouter, Route, Routes} from "react-router";
import {Navigate} from "react-router-dom";
import * as React from "react";
import SignIn from "./pages/sign-in/SignIn";
import SignUp from "./pages/sign-up/SignUp";
import Posts from "./pages/posts/Posts.tsx";

const nvItems = [
    {name: "Posts", path: "/posts"}
];

function Guard({children}: { children: React.ReactNode }) {
    const token = localStorage.getItem("token")
    if (token != null || token != "") {
        return <>{children}</>
    } else {
        return <Navigate to="/sign-in" replace/>
    }
}

interface RouteProps {
    path: string,
    element: React.ReactNode
}

const routes: RouteProps[] = [
    {
        path: "/",
        element: <Guard><Posts/></Guard>
    },
    {
        path: "/posts",
        element: <Guard><Posts/></Guard>
    },
    {
        path: "/sign-in",
        element: <SignIn/>
    },
    {
        path: "/sign-up",
        element: <SignUp/>
    }
]

function App() {
    return (
        <BrowserRouter>
            <Routes>
                {routes.map((route) => (
                    <Route key={route.path} path={route.path} element={route.element}/>
                ))}
            </Routes>
        </BrowserRouter>
    )
}

export default App
