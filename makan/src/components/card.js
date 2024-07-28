// import { ReactComponent as MyIcon } from './path/to/my-icon.svg';

import TagEmoji from './tags';

function Card({
  title,
  distance,
  reviews: { food, place },
  location: { googleMaps },
  tags,
  dishes,
}) {
  return (
    <div className="block max-w-sm p-6 bg-white border border-gray-200 rounded-lg dark:bg-gray-800 dark:border-gray-700 h-screen">
      <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
        {title}
      </h5>
      <div class="flex items-center justify-center mt-5 mb-5">
        <a
          href={googleMaps}
          rel="noreferrer"
          target="_blank"
          class="text-sm text-gray-600 flex items-center"
        >
          <Direction distance={distance} />
        </a>
      </div>
      <div className="flex items-center justify-center space-x-4  mt-5 mb-5">
        {tags.map((tag, i) => (
          <TagEmoji tag={tag} key={i} />
        ))}
      </div>

      {food !== '' && place !== '' ? (
        <ReviewSection food={food} place={place} />
      ) : (
        <BeTheFirstOneToCome dishes={dishes} />
      )}
    </div>
  );
}

function ReviewSection({ food, place }) {
  return (
    <>
      <h4 class="mb-2 text-lg font-bold tracking-tight text-gray-900 dark:text-white">
        Food & Place
      </h4>
      <p class="font-normal text-gray-700 dark:text-gray-400">
        <div dangerouslySetInnerHTML={{ __html: food }} />
      </p>
      <br />
      <p class="font-normal text-gray-700 dark:text-gray-400">
        <div dangerouslySetInnerHTML={{ __html: place }} />
      </p>
    </>
  );
}
function BeTheFirstOneToCome({ dishes }) {
  return (
    <>
      <h4 class="mb-2 text-lg font-bold tracking-tight text-gray-900 dark:text-white">
        Be the first one to try
      </h4>
      <p class="font-normal text-gray-700 dark:text-gray-400">{dishes.join(", ")}</p>
    </>
  );
}

function Direction({ distance }) {
  return (
    <div className="flex">
      <svg
        width="38"
        height="38"
        viewBox="0 0 24 24"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M18.9762 5.5914L14.6089 18.6932C14.4726 19.1023 13.8939 19.1023 13.7575 18.6932L11.7868 12.7808C11.6974 12.5129 11.4871 12.3026 11.2192 12.2132L5.30683 10.2425C4.89772 10.1061 4.89772 9.52743 5.30683 9.39106L18.4086 5.0238C18.7594 4.90687 19.0931 5.24061 18.9762 5.5914Z"
          stroke="#0E46A3"
          stroke-linecap="round"
          stroke-linejoin="round"
        />
      </svg>
      <span class="relative flex h-3 w-3 ">
        <span className="-ml-3">
          <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-sky-400 opacity-75"></span>
          <span class="relative inline-flex rounded-full h-3 w-3 bg-sky-500"></span>
        </span>
      </span>
      <span className="text-sm text-gray-600 -ml-4 mt-3">{distance}m</span>
    </div>
  );
}

export default Card;
