import { IAuth } from "../features/auth/types";
import axiosInstance from "../lib/axiosInstances";

export const login = ({
  username,
  password,
}: Pick<IAuth, "username" | "password">) => {
  return axiosInstance.post("/login", { username, password });
};

export const register = ({
  username,
  password,
}: Pick<IAuth, "username" | "password">) => {
  return axiosInstance.post("/register", {
    username,
    password,
  });
};

export const sessionLogin = ({ token }: { token: string }) => {
  return axiosInstance.get("/session", {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};
