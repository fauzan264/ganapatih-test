"use client";
import Button from "@/components/ui/Button";
import FormInputTextArea from "@/components/ui/FormTextArea";
import { createFeedSchema } from "@/features/app/feed/schemas/createFeedSchema";
import { postFeed } from "@/services/feed";
import useAuthStore from "@/store/useAuthStore";
import { ErrorResponse } from "@/types/error";
import { AxiosError } from "axios";
import { useFormik } from "formik";
import { useRouter } from "next/navigation";
import { toast } from "react-toastify";

export default function CreatePostPage() {
  const router = useRouter();
  const { token } = useAuthStore();

  const onRegister = async ({
    token,
    content,
  }: {
    token: string;
    content: string;
  }) => {
    try {
      const response = await postFeed({
        token,
        content,
      });

      toast.info(response.data.message);
      router.push("/app");
    } catch (error: unknown) {
      const err = error as AxiosError<ErrorResponse>;
      if (err.response) {
        toast.error(err.response.data.message);
      }
    }
  };

  const formik = useFormik({
    initialValues: {
      content: "",
    },
    validationSchema: createFeedSchema,
    onSubmit: ({ content }) => {
      onRegister({
        token,
        content,
      });
    },
  });

  return (
    <div className="max-w-xl mx-auto p-6 bg-base-100 shadow-xl rounded-lg my-10">
      <h1 className="text-3xl font-bold mb-6">Create New Post</h1>
      <form onSubmit={formik.handleSubmit}>
        <FormInputTextArea formik={formik} name="content" label="Content" />
        <Button type="submit" name="Post"></Button>
      </form>
    </div>
  );
}
