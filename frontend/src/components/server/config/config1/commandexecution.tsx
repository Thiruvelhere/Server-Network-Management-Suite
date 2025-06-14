import React, { useState } from 'react';
import './commandexecution.css';
import { FaTerminal, FaPlay } from 'react-icons/fa';
import ModalWrapper from './modalwrapper';

const CommandExecution: React.FC = () => {
  const [showModal, setShowModal] = useState(false);
  const [command, setCommand] = useState('');

  const handleExecute = () => {
    console.log("Command to execute:", command);
    setShowModal(false);
    setCommand('');
  };

  return (
    <>
      <div className="card command-card" onClick={() => setShowModal(true)}>
        <div className="command-icon">
          <FaTerminal size={20} color="white" />
        </div>
        <h3>Command Execution</h3>
        <p>Execute administrative commands remotely</p>
      </div>

      {showModal && (
        <ModalWrapper title="Remote Command Execution" onClose={() => setShowModal(false)}>
          <p className="subtitle">Execute administrative commands remotely with elevated privileges</p>
          <label>Command Input</label>
          <textarea
            placeholder="Enter administrative command..."
            value={command}
            onChange={(e) => setCommand(e.target.value)}
          />
          <button onClick={handleExecute}><FaPlay /> Execute Command</button>
        </ModalWrapper>
      )}
    </>
  );
};

export default CommandExecution;
