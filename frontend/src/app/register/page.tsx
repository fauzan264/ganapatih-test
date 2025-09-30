"use client";
import Button from "@/components/ui/Button";
import FormInput from "@/components/ui/FormInput";
import { registerSchema } from "@/features/auth/register/schemas/registerSchema";
import { register } from "@/services/auth";
import { ErrorResponse } from "@/types/error";
import { AxiosError } from "axios";
import { useFormik } from "formik";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { toast } from "react-toastify";

export default function RegisterPage() {
  const router = useRouter();

  const onRegister = async ({
    username,
    password,
  }: {
    username: string;
    password: string;
  }) => {
    try {
      const response = await register({
        username,
        password,
      });
      toast.info(response.data.message);
      router.push("/");
    } catch (error: unknown) {
      const err = error as AxiosError<ErrorResponse>;
      if (err.response) {
        toast.error(err.response.data.message);
      }
    }
  };

  const formik = useFormik({
    initialValues: {
      username: "",
      password: "",
    },
    validationSchema: registerSchema,
    onSubmit: ({ username, password }) => {
      onRegister({
        username,
        password,
      });
    },
  });
  return (
    <div className="min-h-screen flex items-center justify-center mt-15">
      <div className="card w-4/5 md:w-2/5 card-border card-md bg-slate-50 shadow-sm p-5 rounded-xl">
        <div className="card-body">
          <div className="card-title justify-center text-slate-800">
            Register
          </div>
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
              <Button type="submit" name="Register" />
            </div>
          </form>
          <p className="mt-3 text-slate-800">
            You have an account?{" "}
            <Link href="/" className="text-slate-600">
              Login
            </Link>
          </p>
        </div>
      </div>
    </div>
  );
}
