import {BrowserRouter, Route, Routes} from "react-router";
import * as React from "react";
import SignIn from "./pages/sign-in/SignIn.tsx";

interface RouteProps {
    path: string,
    element: React.ReactNode
}

const routes: RouteProps[] = [
    {
        path: "/",
        element: <div>Home</div>
    },
    {
        path: "/sign-in",
        element: <SignIn/>
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
