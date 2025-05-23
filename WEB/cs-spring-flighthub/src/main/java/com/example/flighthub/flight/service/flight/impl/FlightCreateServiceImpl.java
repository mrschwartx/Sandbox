package com.example.flighthub.flight.service.flight.impl;

import com.example.flighthub.flight.exception.AirportNotFoundException;
import com.example.flighthub.flight.model.Flight;
import com.example.flighthub.flight.model.dto.request.flight.CreateFlightRequest;
import com.example.flighthub.flight.model.entity.AirportEntity;
import com.example.flighthub.flight.model.entity.FlightEntity;
import com.example.flighthub.flight.model.mapper.flight.CreateFlightRequestToFlightEntityMapper;
import com.example.flighthub.flight.model.mapper.flight.FlightEntityToFlightMapper;
import com.example.flighthub.flight.repository.AirportRepository;
import com.example.flighthub.flight.repository.FlightRepository;
import com.example.flighthub.flight.service.flight.FlightCreateService;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;


/**
 * Service interface for creating a flight in the system.
 */
@Service
@RequiredArgsConstructor
public class FlightCreateServiceImpl implements FlightCreateService {

    private final FlightRepository flightRepository;
    private final AirportRepository airportRepository;
    private final CreateFlightRequestToFlightEntityMapper createFlightRequestToFlightEntityMapper =
            CreateFlightRequestToFlightEntityMapper.initialize();
    private final FlightEntityToFlightMapper flightEntityToFlightMapper =
            FlightEntityToFlightMapper.initialize();

    /**
     * Creates a new flight in the system.
     *
     * @param createFlightRequest the request object containing the details of the flight to be created.
     * @return the created {@link Flight} entity.
     */
    @Override
    public Flight createFlight(CreateFlightRequest createFlightRequest) {
        AirportEntity departureAirport = airportRepository.findById(createFlightRequest.getFromAirportId())
                .orElseThrow(() -> new AirportNotFoundException("Departure airport not found with id " + createFlightRequest.getFromAirportId()));
        AirportEntity arrivalAirport = airportRepository.findById(createFlightRequest.getToAirportId())
                .orElseThrow(() -> new AirportNotFoundException("Arrival airport not found with id " + createFlightRequest.getToAirportId()));
        FlightEntity flightEntityTobeSaved = createFlightRequestToFlightEntityMapper.mapForSaving(createFlightRequest, departureAirport, arrivalAirport);
        FlightEntity savedFlight = flightRepository.save(flightEntityTobeSaved);

        return flightEntityToFlightMapper.map(savedFlight);
    }

}
