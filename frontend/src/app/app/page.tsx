"use client";
import FeedPostCard from "@/features/app/feed/components/Feed";
import { IFeed } from "@/features/app/feed/types";
import { getFeeds } from "@/services/feed";
import useAuthStore from "@/store/useAuthStore";
import Link from "next/link";
import { useEffect, useState } from "react";
import { FaPlus } from "react-icons/fa";

export default function AppPage() {
  const [feeds, setFeeds] = useState<IFeed[] | []>([]);
  const [page, setPage] = useState<number>(1);
  const [isLoading, setIsLoading] = useState(true);
  const { token } = useAuthStore();

  const onGetFeeds = async ({
    token,
    page,
    limit,
  }: {
    token: string;
    page: number;
    limit: number;
  }) => {
    try {
      const response = await getFeeds({ token, page, limit });
      setPage(response.data.data.page);
      setFeeds(response.data.data.posts);
    } catch (error) {
      console.error("Failed to load feeds:", error);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    if (token) {
      setIsLoading(true);
      onGetFeeds({ token, page: 1, limit: 10 });
    }
  }, [token]);

  if (isLoading) {
    return (
      <div className="min-h-screen bg-base-200 flex items-center justify-center">
        <span className="loading loading-spinner loading-lg"></span>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-base-200 py-10 pt-25">
      <div className="container mx-auto px-4 max-w-xl">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-3xl font-semibold text-neutral">Feed Post</h1>

          <Link
            href="/app/create"
            className="btn btn-sm bg-slate-700 hover:bg-slate-800 text-slate-100 shadow-lg"
          >
            <FaPlus className="w-4 h-4" />
            Create Post
          </Link>
        </div>

        <div className="text-center mb-8">
          <span className="badge badge-lg badge-neutral font-normal">
            Page {page}
          </span>
        </div>

        {/* List Posts */}
        <div className="space-y-6">
          {feeds && feeds.length > 0 ? (
            feeds.map((post) => (
              <FeedPostCard
                key={post.id}
                token={token}
                userid={post.userid}
                content={post.content}
                createdat={post.createdat}
                id={post.id}
              />
            ))
          ) : (
            <div className="text-center py-10 text-gray-500">
              There are no posts to display yet
            </div>
          )}
        </div>

        <div className="text-center mt-10 text-gray-500 text-sm">
          Menampilkan {feeds.length ?? 0} Post
        </div>
      </div>
    </div>
  );
}
