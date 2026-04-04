import { useNavigate } from "react-router-dom";
import { Button } from "./ui/button";
import { CreateComponent } from "./CreateComponent";

export function HomeComponent() {
  const navigate = useNavigate();

  return (
    <>
      <div>
        <Button onClick={() => navigate("/about")}>About</Button>
        <CreateComponent />
      </div>
    </>
  );
}
