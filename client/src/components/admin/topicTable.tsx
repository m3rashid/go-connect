import React from "react";
import { useSelector } from "react-redux";
import moment from "moment";

import useTopic from "../../hooks/useTopic";
import { Table } from "./table";
import Dialog from "./dialog";

const TopicTable = () => {
  const { deleteTopic, updateTopic } = useTopic();

  const topics = useSelector((state: any) => state.auth.topics);
  const columns = React.useMemo(
    () => [
      {
        Header: "Topic",
        accessor: "name",
        Cell: ({ cell }: { cell: any }) => <>{cell.row.original.name}</>,
      },
      {
        Header: "No. of Posts",
        accessor: "postCount",
        Cell: ({ cell }: { cell: any }) => <>{cell.row.original.postCount}</>,
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
            topicName={cell.row.original.name}
            topicID={cell.row.original.topicID}
          />
        ),
      },
    ],
    // eslint-disable-next-line react-hooks/exhaustive-deps
    []
  );

  const ActionButtons = ({
    topicName,
    topicID,
  }: {
    topicName: string;
    topicID: string;
  }) => (
    <div className="flex gap-2 items-center">
      <Dialog
        isDeleteDialog={true}
        title="Delete Topic"
        content="Are you sure you want to delete this topic?"
        onConfirm={() => deleteTopic(topicID)}
      />
      <Dialog
        isDeleteDialog={false}
        title="Edit Topic"
        content={topicName}
        onConfirm={(newTopicName) => updateTopic(newTopicName || "", topicID)}
      />
    </div>
  );

  return <Table columns={columns} data={topics} title="Topics" />;
};

export default TopicTable;
