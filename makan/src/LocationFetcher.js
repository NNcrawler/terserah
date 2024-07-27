import React from 'react';

function LocationFetcher({
  location: { loading, error, latitude, longitude },
}) {
  if (error) {
    return <p>Enable permissions to access your location data</p>;
  }

  if (loading) {
    return <p>loading... (you may need to enable permissions)</p>;
  }

  return (
    <div>
      <h1 className="text-2xl font-bold underline">Your Location</h1>

      {latitude && (
        <p>
          Latitude: {latitude}, Longitude: {longitude}
        </p>
      )}
      {error && <p>Error: {error}</p>}
    </div>
  );
}

export default LocationFetcher;
