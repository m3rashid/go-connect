import React from "react";

import moment from "moment";
import { useSelector } from "react-redux";

import Dialog from "./dialog";
import { Table } from "./table";
import usePost from "../../hooks/usePost";


const PostsTable = () => {
  const { deletePost } = usePost();
  const posts = useSelector((state: any) => state.auth.posts);
  const columns = React.useMemo(
    () => [
      {
        Header: "Title",
        accessor: "title",
        Cell: ({ cell }: { cell: any }) => (
          <>{cell.row.original.title.substring(0, 30)}</>
        ),
      },
      {
        Header: "Body",
        accessor: "description",
        Cell: ({ cell }: { cell: any }) => (
          <>{cell.row.original.description.substring(0, 30)}</>
        ),
      },
      {
        Header: "Likes",
        accessor: "likes",
        Cell: ({ cell }: { cell: any }) => <>{cell.row.original.likes}</>,
      },
      {
        Header: "Comments",
        accessor: "commentsCount",
        Cell: ({ cell }: { cell: any }) => (
          <>{cell.row.original.commentsCount}</>
        ),
      },
      {
        Header: "Created",
        accessor: "createdAt",
        Cell: ({ cell }: { cell: any }) => (
          <>{moment(cell.row.original.createdAt).format("lll")}</>
        ),
      },
      {
        Header: "Updated",
        accessor: "updatedAt",
        Cell: ({ cell }: { cell: any }) => (
          <>{moment(cell.row.original.updatedAt).format("lll")}</>
        ),
      },
      {
        Header: "Actions",
        accessor: "",
        Cell: ({ cell }: { cell: any }) => (
          <ActionButtons postID={cell.row.original.postID} />
        ),
      },
    ],
    // eslint-disable-next-line react-hooks/exhaustive-deps
    []
  );

  const ActionButtons = ({ postID }: { postID: string }) => (
    <div className="flex gap-2 items-center">
      <Dialog
        isDeleteDialog={true}
        title="Delete Post"
        content="Are you sure you want to delete this post?"
        onConfirm={() => deletePost(postID)}
      />
    </div>
  );

  return <Table columns={columns} data={posts} title="Posts" />;
};

export default PostsTable;
