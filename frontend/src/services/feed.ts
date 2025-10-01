import axiosInstance from "@/lib/axiosInstances";

export const getFeeds = ({
  token,
  page,
  limit,
}: {
  token: string;
  page: number;
  limit: number;
}) => {
  return axiosInstance.get("/feed", {
    params: {
      page,
      limit,
    },
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};

export const postFeed = ({
  token,
  content,
}: {
  token: string;
  content: string;
}) => {
  return axiosInstance.post(
    "/posts",
    { content },
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );
};
