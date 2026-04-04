import { useRoutes } from "react-router-dom";
import "./App.css";
import { routes } from "./routes";

function AppRoutes() {
  return useRoutes(routes);
}

function App() {
  return <AppRoutes />;
}

export default App;
