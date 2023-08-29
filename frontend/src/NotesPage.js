import React, { useState, useEffect } from 'react';

function NotesPage({ user }) {
    const [notes, setNotes] = useState([]);
    const [content, setContent] = useState('');

    useEffect(() => {
        const fetchNotes = async () => {
            const response = await fetch("/notes", {
                headers: {
                    "UserID": user.id
                }
            });

            if (response.ok) {
                const data = await response.json();
                setNotes(data);
            }
        };

        fetchNotes();
    }, [user]);

    const handleAddNote = async (e) => {
        e.preventDefault();
        const response = await fetch("/notes/add", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "UserID": user.id
            },
            body: JSON.stringify({ content })
        });

        if (response.ok) {
            const newNote = await response.json();
            setNotes([...notes, newNote]);
            setContent('');
        } else {
            alert("Failed to add note");
        }
    };

    return (
        <div>
            <h2>Notes</h2>
            <ul>
                {notes.map(note => (
                    <li key={note.id}>{note.content}</li>
                ))}
            </ul>
            <form onSubmit={handleAddNote}>
                <textarea
                    value={content}
                    onChange={e => setContent(e.target.value)}
                    placeholder="New note..."
                />
                <button type="submit">Add Note</button>
            </form>
        </div>
    );
}

export default NotesPage;
