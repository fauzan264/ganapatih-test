"use client";
import Button from "@/components/ui/Button";
import FormInput from "@/components/ui/FormInput";
import { loginSchema } from "@/features/auth/login/schemas/loginSchema";
import { login } from "@/services/auth";
import useAuthStore from "@/store/useAuthStore";
import { ErrorResponse } from "@/types/error";
import { AxiosError } from "axios";
import { useFormik } from "formik";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { toast } from "react-toastify";

export default function Home() {
  const router = useRouter();
  const { setAuth } = useAuthStore();

  const onLogin = async ({
    username,
    password,
  }: {
    username: string;
    password: string;
  }) => {
    try {
      const response = await login({ username, password });
      if (response.status == 200) {
        setAuth({
          token: response.data.token,
          id: "",
          username: "",
        });

        toast.info("Login user successful");
        router.push("/app");
      }
    } catch (error: unknown) {
      const err = error as AxiosError<ErrorResponse>;
      if (err.response) {
        console.log(err.response);
        toast.error(err.response.data.message);
      }
    }
  };

  const formik = useFormik({
    initialValues: {
      username: "",
      password: "",
    },
    validationSchema: loginSchema,
    onSubmit: ({ username, password }) => {
      onLogin({ username, password });
    },
  });

  return (
    <div className="min-h-screen flex items-center justify-center mt-15">
      <div className="card w-4/5 md:w-2/5 card-border card-md bg-slate-50 shadow-sm p-5 rounded-xl">
        <div className="card-body">
          <div className="card-title justify-center text-slate-800">Login</div>
          <form onSubmit={formik.handleSubmit}>
            <div className="flex flex-wrap">
              <FormInput
                formik={formik}
                name="username"
                label="Username"
                type="text"
              />
              <FormInput
                formik={formik}
                name="password"
                label="Password"
                type="password"
              />
              <Button type="submit" name="Login" />
            </div>
          </form>
          <p className="mt-3 text-slate-800 mb-5">
            You have an account?{" "}
            <Link href="/register" className="text-slate-600">
              Register
            </Link>
          </p>
        </div>
      </div>
    </div>
  );
}
