import React from "react";
import { useState } from "react";
import { shortUrl } from "./api.ts";
import "./App.css";

function App() {
  const [longUrl, setLongUrl] = useState("");
  const [userId, setUserId] = useState("");
  const [shortenedUrl, setShortenedUrl] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    console.log("Submit button clicked");
    setShortenedUrl(null);
    setError(null);
    if (!longUrl || !userId) {
      setError("Enter both Url and Id");
      return;
    }
    const response = await shortUrl({ LongUrl: longUrl, UserId: userId });
    console.log(response);
    if (response) {
      setShortenedUrl(response.short_url);
    } else {
      setError("Failed to Create URL");
    }
  };
  return (
    <>
      <div>
        <h1>URL Shortener</h1>
      </div>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="Enter URL to be shortened"
          value={longUrl}
          onChange={(e) => setLongUrl(e.target.value)}
        />
        <br />
        <input
          type="text"
          placeholder="Enter user ID"
          value={userId}
          onChange={(e) => setUserId(e.target.value)}
        />
        <br />
        <button type="submit" onClick={()=>console.log("Clicked")} style={{ padding: "10px 20px" }}>Shorten URL</button>
      </form>
      {shortenedUrl && (
        <p>
          Shortened URL : <a href={shortenedUrl}>{shortenedUrl}</a>
        </p>
      )}
      {error && <p style={{ color: "red" }}>{error}</p>}
    </>
  );
}

export default App;
