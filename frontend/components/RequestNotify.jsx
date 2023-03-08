import React, { useEffect, useState } from "react"

export const RequestNotify = ({ type, accepted, rejected }) => {
    const [message, setMessage] = useState("")
    const handleAccepted = () => {
        accepted();
    }
    const handleRejected = () => {
        rejected();
    }
    useEffect(() => {
        if (type.hasOwnProperty("group-id")) {
            let str = `${type["admin"]}` + " Would Like You To Join " + `${type["group-name"]}`
            setMessage(str)
        } else if (type.hasOwnProperty("followRequest")) {
            let str = `${type["followRequest-username"]}` + " Would Like To Follow You"
            setMessage(str)
        }
    }, [type])

    return (
        <div>
            <p>{message}</p>
            <button className="button-85" onClick={handleAccepted}>Accept</button>
            {/* <button className="button-85" onClick={handleRejected}>Reject</button> */}
        </div>
    )
}

export const GroupPostNotify = ({ type }) => {
    const [message, setMessage] = useState("")
    useEffect(() => {
        if (type.hasOwnProperty("group-id")) {
            let str = `${type["author"]}` + " Added Post To " + `${type["group"]["group-name"]}`
            setMessage(str)
        }
    }, [type])
    const handleBrokenAuthImage = (source) => {
        if (source != "") {
            return source
        } else {
            return "https://www.transparentpng.com/thumb/user/gray-user-profile-icon-png-fP8Q1P.png"
        }
    }

    return (
        <div>
            <p>{message}</p>
            <button type="button" className="alert-group-post-button" disabled={true} >
                <img src={handleBrokenAuthImage(type["group"]["group-avatar"])} />
                <p className="group-post-button-name">{type["group"]["group-name"]}</p>
            </button>
        </div>
    )
}

export const RemoveGroupNotify = ({ type }) => {
    const [message, setMessage] = useState("")
    useEffect(() => {
        if (type.hasOwnProperty("group-id")) {
            let str = "🙃 You Have Been Removed From " + `${type["group-name"]}`
            setMessage(str)
        }
    }, [type])
    const handleBrokenAuthImage = (source) => {
        if (source != "") {
            return source
        } else {
            return "https://www.transparentpng.com/thumb/user/gray-user-profile-icon-png-fP8Q1P.png"
        }
    }

    return (
        <div>
            <p>{message}</p>
            <button type="button" className="alert-group-post-button" disabled={true} >
                <img src={handleBrokenAuthImage(type["group-avatar"])} />
                <p className="group-post-button-name">{type["group-name"]}</p>
            </button>
        </div>
    )
}

export const AddedGroupNotify = ({ type }) => {
    const [message, setMessage] = useState("")
    useEffect(() => {
        let str
        if (type["action"]) {
            str = "😃" + `${type["user"]}` + " Has Now Joined " + `${type["group-name"]}`
            // setMessage(str)
        } else {
            str = "🙃 You Have Been Removed From " + `${type["group-name"]}`
        }
        setMessage(str)
    }, [type])
    const handleBrokenAuthImage = (source) => {
        if (source != "") {
            return source
        } else {
            return "https://www.transparentpng.com/thumb/user/gray-user-profile-icon-png-fP8Q1P.png"
        }
    }

    return (
        <div>
            <p>{message}</p>
            <button type="button" className="alert-group-post-button" disabled={true} >
                <img src={handleBrokenAuthImage(type["group-avatar"])} />
                <p className="group-post-button-name">{type["group-name"]}</p>
            </button>
        </div>
    )
}