function simulateAPICall() {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      const response = {
        data: [
          {
            name: 'Bakso Join Blok S',
            dishType: 'bakso',
            location: {
              googleMaps: 'https://goo.gl/maps/1J1J1J1J1J1J1J1J1',
              address: 'Jl. Blok S No. 1, Jakarta Selatan',
            },
          },
          {
            name: 'Nasi Goreng Kambing',
            dishType: 'nasi goreng',
            location: {
              googleMaps: 'https://goo.gl/maps/2J2J2J2J2J2J2J2J2',
              address: 'Jl. Kambing No. 2, Jakarta Selatan',
            },
          },
          {
            name: 'Sate Ayam Madura',
            dishType: 'sate',
            location: {
              googleMaps: 'https://goo.gl/maps/3J3J3J3J3J3J3J3J3',
              address: 'Jl. Madura No. 3, Jakarta Selatan',
            },
          },
        ],
      };
      resolve(response);
    }, 30); // Simulating a delay of 1 second
  });
}
