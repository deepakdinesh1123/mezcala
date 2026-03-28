import { BrowserRouter, Outlet, Route, Routes } from "react-router-dom";
import "./App.css";
import { MenuComponent } from "./components/MenuComponent";
import AboutComponent from "./components/AboutComponent";
function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<MenuComponent />} />
        <Route path="/about" element={<AboutComponent />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
