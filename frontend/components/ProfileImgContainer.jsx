import React from "react";

export default function ProfileImgContainer(props) {
  return (
    <div>
      {props.name ? (
        <div>
          <div className="profileImgContainer">
            <div className="profileImgBg"></div>
            <img className="profileImg" src={props.avatar} alt="" />
            <span className="firstLast">
              {props.name} {props.user.last}
            </span>
            <p className="aboutme">{props.user.aboutme}</p>
            <hr className="break" />

            <div className="followerDiv">
              <div>
                <span className="followerFollowing">Following</span>
              </div>
              <div>
                <span className="count">34</span>
              </div>
            </div>
            <hr className="break" />

            <div className="followerDiv">
              <div>
                <span className="followerFollowing">Followers</span>
              </div>
              <div>
                <span className="count">155</span>
              </div>
            </div>
            <hr className="break" />
            <div className="followerDiv">
              <span className="redText">View Profile</span>
            </div>
          </div>
        </div>
      ) : (
        <div> loading... </div>
      )}
    </div>
  );
}