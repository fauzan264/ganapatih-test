import { create } from "zustand";
import { persist } from "zustand/middleware";
import useFollowStore from "./useFollowStore";

interface IUseAuthStoreState {
  id: string;
  username: string;
  token: string;
}

interface IUseAuthStore extends IUseAuthStoreState {
  setAuth: ({ token, id, username }: IUseAuthStoreState) => void;
  logout: () => void;
}

const useAuthStore = create<IUseAuthStore>()(
  persist(
    (set) => ({
      token: "",
      id: "",
      username: "",
      setAuth: ({
        token,
        id,
        username,
      }: {
        token: string;
        id: string;
        username: string;
      }) => {
        set({ token, id, username });
      },
      logout: () => {
        // Clear auth state
        set({ token: "", id: "", username: "" });
        localStorage.removeItem("authToken");

        // Clear follow state
        const followStore = useFollowStore.getState();
        followStore.clearFollows();
      },
    }),
    {
      name: "authToken",
      partialize: (state) => ({ token: state.token }),
    }
  )
);

export default useAuthStore;
