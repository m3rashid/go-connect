import moment from "moment";
import React from "react";
import Avatar from "react-nice-avatar";
import { Link } from "react-router-dom";
import { IAvatar } from "../../interfaces/avatar";

import Loader from "../loader";

interface INotif {
  notif?: any;
  bgColor?: string;
  isNotification?: boolean;
}

const Notif: React.FC<INotif> = ({ notif, bgColor, isNotification = true }) => {
  if (!notif.avatarID) {
    return <Loader />;
  }

  const avatar: IAvatar = {
    avatarID: notif.avatarID,
    sex: notif.sex,
    faceColor: notif.faceColor,
    earSize: notif.earSize,
    hairColor: notif.hairColor,
    hairStyle: notif.hairStyle,
    hatColor: notif.hatColor,
    hatStyle: notif.hatStyle,
    glassesStyle: notif.glassesStyle,
    noseStyle: notif.noseStyle,
    mouthStyle: notif.mouthStyle,
    shirtStyle: notif.shirtStyle,
    shirtColor: notif.shirtColor,
    bgColor: notif.bgColor,
    isGradient: notif.isGradient,
  };

  return (
    <>
      <div className="dark:text-gray-200 pl-2 pb-2 flex gap-4 w-full">
        <div className="">
          <Avatar className="h-16 w-16" {...avatar} />
        </div>
        <div
          className={`flex flex-col items-start justify-center bg-gray-50 dark:${
            bgColor || "bg-gray-700"
          } w-full px-3 py-2 rounded-md`}
        >
          <div className="font-semibold">
            <Link to={`/user/${notif.userId || notif.userID}`}>
              @{notif.userName}
            </Link>{" "}
            &nbsp;
            <span className="font-normal">
              {!isNotification ? "on " : "at "}
              {moment(notif.comment_createdAt).format("DD/MM/YYYY hh:mm A")}
            </span>
          </div>
          <div className={isNotification ? "my-1" : ""}>{notif.text}</div>
        </div>
      </div>
    </>
  );
};

export default Notif;
