"use client";

import { sessionLogin } from "@/services/auth";
import useAuthStore from "@/store/useAuthStore";
import { useEffect } from "react";
import { toast } from "react-toastify";

export default function AuthProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const { token, setAuth } = useAuthStore();

  const onAuthSessionLogin = async () => {
    try {
      if (!token) return;
      const response = await sessionLogin({ token });

      setAuth({
        token,
        id: response?.data?.data?.id,
        username: response?.data?.data?.username,
      });
    } catch (error) {
      toast.error(`session login error: ${error}`);
    }
  };

  useEffect(() => {
    if (token) {
      onAuthSessionLogin();
    }
  }, [token]);

  return <>{children}</>;
}
