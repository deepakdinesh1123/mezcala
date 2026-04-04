import { MenuComponent } from "@/components/MenuComponent";
import { Outlet } from "react-router-dom";

const MainLayout: React.FC = () => {
  return (
    <>
      <div className="sticky top-0 z-10">
        <MenuComponent />
      </div>
      <main>
        <Outlet />
      </main>
    </>
  );
};

export default MainLayout;
