import { useEffect, useState } from "react";

interface Engine {
  name: string;
  verion: string;
}

export const useEngines = () => {
  const [engines, setEngines] = useState<Engine[]>([{ name: "", verion: "" }]);

  useEffect(() => {
    setEngines([
      { name: "PostgreSQL", verion: "16" },
      { name: "MySQL", verion: "1.7" },
    ]);
  }, []);

  return { engines };
};
