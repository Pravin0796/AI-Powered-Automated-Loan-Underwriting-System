import React, { useState } from "react";
import { loanClient } from "../services/grpc";
import { LoanRequest } from "../proto/loan";

const ApplyLoan: React.FC = () => {
    const [userId, setUserId] = useState(1);
    const [amount, setAmount] = useState(1000);
    const [term, setTerm] = useState(12);
    const [loanId, setLoanId] = useState<number | null>(null);
    const [status, setStatus] = useState<string>("");

    const handleApply = async () => {
        const request: LoanRequest = {
            userId,
            amountRequested: amount,
            loanTerm: term,
        };

        const res = await loanClient.ApplyForLoan(request);
        setLoanId(res.loanId);
        setStatus(res.status);
    };

    return (
        <div>
            <h2>Apply for Loan</h2>
            <input type="number" value={userId} onChange={(e) => setUserId(Number(e.target.value))} placeholder="User ID" />
            <input type="number" value={amount} onChange={(e) => setAmount(Number(e.target.value))} placeholder="Amount" />
            <input type="number" value={term} onChange={(e) => setTerm(Number(e.target.value))} placeholder="Term (months)" />
            <button onClick={handleApply}>Apply</button>
            {loanId && <p>Loan ID: {loanId}, Status: {status}</p>}
        </div>
    );
};

export default ApplyLoan;
