import AboutComponent from "@/components/AboutComponent";
import { HomeComponent } from "@/components/HomeComponent";
import RequireAuth from "@/components/RequireAuth";
import MainLayout from "@/layouts/MainLayout";
import type { RouteObject } from "react-router-dom";

export const routes: RouteObject[] = [
  {
    element: <RequireAuth />,
    children: [
      {
        path: "/",
        element: <MainLayout />,
        children: [
          { index: true, element: <HomeComponent /> },
          { path: "/about", element: <AboutComponent /> },
        ],
      },
    ],
  },
];
