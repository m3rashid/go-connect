import React from "react";
import { useDispatch, useSelector } from "react-redux";
import Loader from "../components/loader";

import { getPosts } from "../store/actions/post.action";
const RightSidebar = React.lazy(() => import("../components/nav/rightSidebar"));
const CreatePost = React.lazy(() => import("../components/user/create"));
const PostCard = React.lazy(() => import("../components/user/postCard"));

const Main = () => {
  const dispatch = useDispatch();
  const auth = useSelector((state: any) => state.auth);
  const loading = useSelector((state: any) => state.posts.postsLoading);
  const user = auth.user;

  React.useEffect(() => {
    dispatch(getPosts(user) as any);
  }, [dispatch, user]);

  const posts = useSelector((state: any) =>
    Object.values(state.posts.posts).sort(
      (a: any, b: any) => b.updatedAt - a.updatedAt
    )
  );

  return (
    <>
      <div className="flex flex-col items-center gap-4  md:w-auto m-[10px]">
        <CreatePost />
        <div className="w-full rightSidebarInMainPage">
          <RightSidebar fullWidth />
        </div>

        {loading ? (
          <Loader />
        ) : (
          posts &&
          posts.map((post: any) => (
            <PostCard key={post.postID} post={post} loggedUser={user} />
          ))
        )}
      </div>
    </>
  );
};

export default Main;
