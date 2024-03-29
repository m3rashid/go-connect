import React from "react";
import { useSelector } from "react-redux";
import moment from "moment";

import { Table } from "./table";
import Dialog from "./dialog";
import useDeleteUser from "../../hooks/useDeleteUser";

const UserTable = () => {
  const { deleteUser } = useDeleteUser();
  const users = useSelector((state: any) => state.auth.users);
  const columns = React.useMemo(
    () => [
      {
        Header: "Username",
        accessor: "userName",
        Cell: ({ cell }: { cell: any }) => <>{cell.row.original.userName}</>,
      },
      {
        Header: "Name",
        accessor: "",
        Cell: ({ cell }: { cell: any }) => (
          <>{cell.row.original.firstName + " " + cell.row.original.lastName}</>
        ),
      },
      {
        Header: "Email",
        accessor: "email",
        Cell: ({ cell }: { cell: any }) => <>{cell.row.original.email}</>,
      },
      {
        Header: "Phone",
        accessor: "phNumber",
        Cell: ({ cell }: { cell: any }) => (
          <>{cell.row.original.phNumber || "--"}</>
        ),
      },
      {
        Header: "Reputation",
        accessor: "reputation",
        Cell: ({ cell }) => <>{cell.row.original.reputation || "--"}</>,
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
          <ActionButtons
            userID={cell.row.original.userID}
            avatarID={cell.row.original.avatarID}
          />
        ),
      },
    ],
    // eslint-disable-next-line react-hooks/exhaustive-deps
    []
  );

  const ActionButtons = ({
    userID,
    avatarID,
  }: {
    userID: string;
    avatarID: string;
  }) => (
    <div className="flex gap-2 items-center">
      <Dialog
        isDeleteDialog={true}
        title="Delete User"
        content="Are you sure you want to delete this user?"
        onConfirm={() => deleteUser(userID, avatarID)}
      />
    </div>
  );

  return <Table columns={columns} data={users} title="Users" />;
};

export default UserTable;
