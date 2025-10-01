import * as yup from "yup";

export const createFeedSchema = yup.object().shape({
  content: yup.string().required("Content is required"),
});
