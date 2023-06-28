import React from "react";
import { Link } from "react-router-dom";

export default function Home() {
  return (
    <div>
      <Link to="/about">About page</Link>
      <Link to="/cabout">Babout page</Link>
    </div>
  );
}
