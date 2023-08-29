import React, { useState } from 'react';
import LoginPage from './LoginPage';
import RegisterPage from './RegisterPage';
import NotesPage from './NotesPage';

function App() {
    const [user, setUser] = useState(null);

    if (!user) {
        return (
            <div>
                <LoginPage onLogin={setUser} />
                <RegisterPage onRegister={setUser} />
            </div>
        );
    }

    return <NotesPage user={user} />;
}

export default App;
