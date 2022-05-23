import axios from "axios";
import { toast } from "react-toastify";

import {
  POSTS_LOADING,
  POSTS_LOADED,
  BOOKMARKS_LOADING,
  BOOKMARKS_LOADED,
  BOOKMARK_ADDED,
  BOOKMARK_DELETED,
  SERVER_ROOT_URL,
  tokenConfig,
  ADD_LIKE,
  REMOVE_LIKE,
} from "../constants";

export const postsLoading = () => ({ type: POSTS_LOADING });
export const bookmarksLoading = () => ({ type: BOOKMARKS_LOADING });

export const getPosts =
  ({ userID }: { userID: string }) =>
  (dispatch: any) => {
    dispatch(postsLoading());
    const body = JSON.stringify({ userID });
    axios
      .post(`${SERVER_ROOT_URL}/post/all`, body, tokenConfig())
      .then((res) => dispatch({ type: POSTS_LOADED, payload: res.data }))
      .catch((err: any) => toast.error("Error loading posts"));
  };

export const getAllBookmarks =
  ({ userID }: { userID: string }) =>
  (dispatch: any) => {
    dispatch(bookmarksLoading());
    const body = JSON.stringify({ userID });
    axios
      .post(`${SERVER_ROOT_URL}/bookmark/all`, body, tokenConfig())
      .then((res) => dispatch({ type: BOOKMARKS_LOADED, payload: res.data }))
      .catch((err: any) => toast.error("Error loading bookmarks"));
  };

export const addBookmark =
  (userID: string, postID: string) => (dispatch: any) => {
    dispatch(bookmarksLoading());
    const body = JSON.stringify({ userID, postID });
    axios
      .post(`${SERVER_ROOT_URL}/bookmark/add`, body, tokenConfig())
      .then((res) => dispatch({ type: BOOKMARK_ADDED, payload: res.data }))
      .catch((err: any) => toast.error("Error adding bookmark"));
  };

export const removeBookmark =
  (userID: string, postID: string) => (dispatch: any) => {
    dispatch(bookmarksLoading());
    const body = JSON.stringify({ userID, postID });
    axios
      .post(`${SERVER_ROOT_URL}/bookmark/remove`, body, tokenConfig())
      .then((res) => dispatch({ type: BOOKMARK_DELETED, payload: res.data }))
      .catch((err: any) => toast.error("Error removing bookmark"));
  };

export const addLike = (userID: string, postID: string) => (dispatch: any) => {
  const body = JSON.stringify({ userID, postID });
  axios
    .post(`${SERVER_ROOT_URL}/like/add`, body, tokenConfig())
    .then((res) => dispatch({ type: ADD_LIKE, payload: { postID } }))
    .catch((err: any) => toast.error("Error adding like"));
};

export const removeLike =
  (userID: string, postID: string) => (dispatch: any) => {
    const body = JSON.stringify({ userID, postID });
    axios
      .post(`${SERVER_ROOT_URL}/like/remove`, body, tokenConfig())
      .then((res) => dispatch({ type: REMOVE_LIKE, payload: { postID } }))
      .catch((err: any) => {
        console.log(err);
        toast.error("Error removing like");
      });
  };
