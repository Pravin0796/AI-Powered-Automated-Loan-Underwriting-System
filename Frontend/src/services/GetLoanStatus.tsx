import React, { useState } from "react";
import { loanClient } from "../services/grpc";

const GetLoanStatus: React.FC = () => {
    const [loanId, setLoanId] = useState<number>(0);
    const [status, setStatus] = useState("");

    const fetchStatus = async () => {
        const res = await loanClient.GetLoanStatus({ loanId });
        setStatus(res.status);
    };

    return (
        <div>
            <h2>Check Loan Status</h2>
            <input type="number" value={loanId} onChange={(e) => setLoanId(Number(e.target.value))} placeholder="Loan ID" />
            <button onClick={fetchStatus}>Get Status</button>
            {status && <p>Status: {status}</p>}
        </div>
    );
};

export default GetLoanStatus;
