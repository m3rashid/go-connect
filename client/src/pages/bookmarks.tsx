import React from "react";
import { useDispatch, useSelector } from "react-redux";

import PostCard from "../components/user/postCard";
import { getAllBookmarks } from "../store/actions/post.action";
import Loader from "../components/loader";

const Bookmarks = () => {
  const dispatch = useDispatch();
  const auth = useSelector((state: any) => state.auth);
  const loading = useSelector((state: any) => state.posts.bookmarksLoading);
  const user = auth.user;

  const [refresh, setRefresh] = React.useState(false);

  React.useEffect(() => {
    window.scrollTo(0, 0);
    dispatch(getAllBookmarks(user) as any);
  }, [dispatch, user, refresh]);

  const Reload = () => {
    setRefresh(!refresh);
  };

  const bookmarks = useSelector((state: any) =>
    Object.values(state.posts.bookmarks).sort(
      (a: any, b: any) => b.updatedAt - a.updatedAt
    )
  );

  return (
    <>
      <div className="flex flex-col items-center gap-4  md:w-auto m-[10px]">
        {loading ? (
          <Loader />
        ) : bookmarks.length > 0 ? (
          bookmarks.map((bookmark: any) => (
            <PostCard
              key={bookmark.bookmarkID}
              post={bookmark}
              loggedUser={user}
              reload={Reload}
            />
          ))
        ) : (
          <h3 className="text-center dark:text-gray-200 font-semibold text-2xl">
            No bookmarks found
          </h3>
        )}
      </div>
    </>
  );
};

export default Bookmarks;
