const FAQ = () => {
    const faqs = [
      {
        question: "How do I apply for a loan?",
        answer: "You can apply for a loan by navigating to the 'Apply Loan' page and filling out the required details.",
      },
      {
        question: "How can I track my loan application?",
        answer: "Go to the 'Loan Status' section to view the progress of your loan application.",
      },
      {
        question: "What is a credit score?",
        answer: "A credit score is a numerical representation of your creditworthiness.",
      },
    ];
  
    return (
      <div className="max-w-4xl mx-auto p-6 bg-white shadow-md rounded-lg">
        <h2 className="text-2xl font-bold mb-4">Frequently Asked Questions</h2>
        {faqs.map((faq, index) => (
          <div key={index} className="mb-4">
            <h3 className="font-semibold">{faq.question}</h3>
            <p className="text-gray-700">{faq.answer}</p>
          </div>
        ))}
      </div>
    );
  };
  
  export default FAQ;