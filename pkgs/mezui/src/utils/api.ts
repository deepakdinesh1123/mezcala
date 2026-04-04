import { Configuration } from "../oas-client/configuration";
import { DefaultApi } from "../oas-client/api";
import axios from "axios";

const config = new Configuration({
  basePath: `${import.meta.env.VITE_BASE_PATH}`,
});

const axiosInstance = axios.create({
  baseURL: config.basePath,
  timeout: 30000,
});

axiosInstance.interceptors.request.use(
  (config) => {
    config.headers = config.headers || {};
    config.headers["Content-Type"] = "application/json";
    return config;
  },
  (error) => Promise.reject(error),
);

axiosInstance.interceptors.response.use(
  (response) => response,
  (error) => Promise.reject(error),
);

export const api = new DefaultApi(config, config.basePath, axiosInstance);
