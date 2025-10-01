"use client";
import { FaUserMinus, FaUserPlus } from "react-icons/fa";
import { IFeed } from "../types";
import { useState } from "react";
import useAuthStore from "@/store/useAuthStore";
import { toast } from "react-toastify";
import useFollowStore from "@/store/useFollowStore";

const formatTimestamp = (isoString: string) => {
  return new Date(isoString).toLocaleString("id-ID", {
    day: "numeric",
    month: "short",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
};

export default function FeedPostCard({
  token,
  userid,
  content,
  createdat,
  id: postId,
}: Omit<IFeed, "id"> & { token: string; id: number }) {
  const [isLoading, setIsLoading] = useState(false);
  const { id: currentUserId } = useAuthStore();

  const followingIds = useFollowStore((state) => state.followingIds);
  const toggleFollowGlobal = useFollowStore((state) => state.toggleFollow);

  const isFollowing = followingIds.includes(userid);

  const handleToggleFollow = async () => {
    if (isLoading) return;

    setIsLoading(true);

    try {
      const result = await toggleFollowGlobal(userid, token, isFollowing);

      if (result.success) {
        toast.success(result.message);
      } else {
        toast.error(result.message);
      }
    } catch (error) {
      toast.error("Failed to connect to server.");
    } finally {
      setIsLoading(false);
    }
  };

  const FollowButton = isFollowing ? (
    <button
      className="btn btn-xs btn-outline btn-warning"
      onClick={handleToggleFollow}
      disabled={isLoading}
    >
      {isLoading ? (
        <span className="loading loading-spinner loading-xs"></span>
      ) : (
        <FaUserMinus className="w-3 h-3" />
      )}
      {isLoading ? "Process..." : "Unfollow"}
    </button>
  ) : (
    <button
      className="btn btn-xs btn-info text-slate-100"
      onClick={handleToggleFollow}
      disabled={isLoading}
    >
      {isLoading ? (
        <span className="loading loading-spinner loading-xs"></span>
      ) : (
        <FaUserPlus className="w-3 h-3" />
      )}
      {isLoading ? "Process..." : "Follow"}
    </button>
  );

  return (
    <div className="card w-full bg-base-100 shadow-xl border border-gray-100 hover:shadow-2xl transition-shadow duration-300">
      <div className="card-body p-5">
        <div className="flex justify-between items-center text-sm text-gray-500 border-b border-dashed border-gray-200 pb-2 mb-3">
          <div className="flex items-center space-x-3">
            <span className="font-semibold text-gray-700">
              User ID:{" "}
              <span className="text-secondary font-extrabold">{userid}</span>
            </span>

            {Number(currentUserId) !== userid && FollowButton}
          </div>

          <span className="text-xs">{formatTimestamp(createdat)}</span>
        </div>

        <p className="text-lg font-medium text-gray-800 leading-relaxed mb-4">
          {content}
        </p>

        <div className="card-actions justify-end">
          <div className="badge badge-outline badge-sm text-gray-400">
            Post ID: {postId}
          </div>
        </div>
      </div>
    </div>
  );
}
