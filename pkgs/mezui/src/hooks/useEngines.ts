import type {
  DatabaseConfig,
  GetSupportedEngines200ResponseInner,
} from "@/oas-client";
import { api } from "@/utils/api";
import { useEffect, useState } from "react";

export const useEngines = () => {
  const [loading, setLoading] = useState(true)
  const [engines, setEngines] = useState<GetSupportedEngines200ResponseInner[]>(
    [],
  );

  useEffect(() => {
    const fetchEngines = async () => {
      try {
        const response = await api.getSupportedEngines();
        setEngines(response.data);
        console.log(response.data);
        
      } catch (e: unknown) {
        console.error("Unable to fetch supported engines:", e);
      } finally {
        setLoading(false)
      }
    };

    fetchEngines();
    
  }, []);

  const createDatabase = async (databaseConfig: DatabaseConfig) => {
    try {
      const response = await api.createDatabase(databaseConfig);
      console.log(response.data);
    } catch (e: unknown) {
      console.error("Unable to create database:", e);
    }
  };

  return { engines, createDatabase, loading };
};
