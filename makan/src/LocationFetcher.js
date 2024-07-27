import React, { useState } from 'react';
import { useGeolocation } from '@uidotdev/usehooks';

function LocationFetcher() {
  const location = useGeolocation();

  if (location.loading) {
    return <p>loading... (you may need to enable permissions)</p>;
  }

  if (location.error) {
    return <p>Enable permissions to access your location data</p>;
  }

  return (
    <div>
      <h1>Get User Location</h1>

      {location.latitude && (
        <p>
          Latitude: {location.latitude}, Longitude: {location.longitude}
        </p>
      )}
      {location.error && <p>Error: {location.error}</p>}
    </div>
  );
}

export default LocationFetcher;
