import {BrowserRouter, Route, Routes} from "react-router";
import {Navigate} from "react-router-dom";
import * as React from "react";
import SignIn from "./pages/sign-in/SignIn";
import SignUp from "./pages/sign-up/SignUp";

interface RouteProps {
    path: string,
    element: React.ReactNode
}

function Guard({children}: { children: React.ReactNode }) {
    const token = localStorage.getItem("token")
    if (token) {
        return children;
    } else {
        return <Navigate to="/sign-in" replace/>
    }
}

const routes: RouteProps[] = [
    {
        path: "/",
        element: <Guard>
            <div>Home</div>
        </Guard>
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
