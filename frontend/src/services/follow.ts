import axiosInstance from "@/lib/axiosInstances";

export const follow = ({
  token,
  userid,
}: {
  token: string;
  userid: number;
}) => {
  return axiosInstance.post(`/follow/${userid}`, null, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};

export const unfollow = ({
  token,
  userid,
}: {
  token: string;
  userid: number;
}) => {
  return axiosInstance.delete(`/follow/${userid}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};
