import { create } from "zustand";
import { persist } from "zustand/middleware";
import { AxiosError } from "axios";
import { follow, unfollow } from "@/services/follow";
import { ErrorResponse } from "@/types/error";

interface IFollowStoreState {
  followingIds: number[];
}

interface IFollowStoreActions {
  toggleFollow: (
    userid: number,
    token: string,
    isCurrentlyFollowing: boolean
  ) => Promise<{ success: boolean; message: string }>;
  setInitialFollows: (ids: number[]) => void;
  clearFollows: () => void;
  addFollowing: (userid: number) => void;
  removeFollowing: (userid: number) => void;
}

interface IUseFollowStore extends IFollowStoreState, IFollowStoreActions {}

const useFollowStore = create<IUseFollowStore>()(
  persist(
    (set) => ({
      followingIds: [],

      setInitialFollows: (ids) => {
        set({ followingIds: ids });
      },

      clearFollows: () => {
        set({ followingIds: [] });
      },

      addFollowing: (userid) => {
        set((state) => ({
          followingIds: [...state.followingIds, userid],
        }));
      },

      removeFollowing: (userid) => {
        set((state) => ({
          followingIds: state.followingIds.filter((id) => id !== userid),
        }));
      },

      toggleFollow: async (userid, token, isCurrentlyFollowing) => {
        try {
          if (isCurrentlyFollowing) {
            const response = await unfollow({ token, userid });

            set((state) => ({
              followingIds: state.followingIds.filter((id) => id !== userid),
            }));

            return {
              success: true,
              message: response.data.message || "Success unfollow",
            };
          } else {
            const response = await follow({ token, userid });

            set((state) => ({
              followingIds: [...state.followingIds, userid],
            }));

            return {
              success: true,
              message: response.data.message || "Success follow",
            };
          }
        } catch (error: unknown) {
          const err = error as AxiosError<ErrorResponse>;
          const message =
            err.response?.data?.message ||
            "A connection or server error occurred.";

          return {
            success: false,
            message,
          };
        }
      },
    }),
    {
      name: "follow-storage",
      partialize: (state) => ({ followingIds: state.followingIds }),
    }
  )
);

export default useFollowStore;
